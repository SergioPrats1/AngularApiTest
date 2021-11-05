import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { throwError } from "rxjs";
import { environment } from '../../environments/environment';
import { AuthenticationService } from '../services';

export interface Song {
  id: number;
  title: string;
  artist: string;
  year: number;
  comments: string;
}

export interface NewSong {
  title: string;
  artist: string;
  year: number;
  comments: string;
}

@Injectable({
  providedIn: 'root'
})
export class ApiReaderService {

  // C# API (legacy)
  //getUrl = 'http://localhost:5000/api/songs';	  
  
  //ip = 'localhost';
  //ip = '192.168.1.110';
  //port = "8080";
  getUrl: string;
  addUrl: string;
  delUrl: string;


  constructor(private http: HttpClient, private authenticationService: AuthenticationService) { 
    // Golang API
    //let baseUrl = "http://" + this.ip + ":" + this.port;
    let baseUrl = environment.apiEndPoint;

    this.getUrl = baseUrl + "/getSongs";
    this.addUrl = baseUrl + "/addSong";
    this.delUrl = baseUrl + "/deleteSong";
  }
 
  //configUrl = 'assets/Songs.json';
  
  getSongs() {   
  
    let options = {
				headers: {
					'Content-Type': 'application/json',
					'Access-Control-Allow-Origin': '*',
          'Authorization': this.authenticationService.token,
				},
				method: 'GET', 
				mode: '*' 
			};  
  
    return this.http.get<Song[]>(this.getUrl, options)
	  .pipe(
        catchError(this.handleError)
	  );

    //return this.http.get(this.getUrl);

  }  
 
  addSong(newSong: NewSong) {

    let options = {
      headers: {
        'Authorization': this.authenticationService.token,
      }
    };  

	  return this.http.post<NewSong>(this.addUrl, JSON.stringify(newSong), options)    
	    .pipe(catchError(this.handleError));
  } 
 
  deleteSong(id: number) {
    let options = {
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*',
            'Authorization': this.authenticationService.token,
          },
          method: 'GET',
          mode: '*'
        };  	  
    
    let url = this.delUrl + '/' + id.toString();

    console.log(url);

    return this.http.get<NewSong>(url, options)    
      .pipe(catchError(this.handleError));
    } 

  private handleError(error: HttpErrorResponse) {
    console.error(error.message);
    return throwError('A data error occurred, please try again.');
  }
 
}
