import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

export interface Song {
  id: number;
  title: string;
  artist: string;
  year: number;
}

export interface NewSong {
  title: string;
  artist: string;
  year: number;
}

@Injectable({
  providedIn: 'root'
})
export class ApiReaderService {

  // C# API (legacy)
  //getUrl = 'http://localhost:5000/api/songs';	  
  
  //ip = 'localhost';
  ip = '192.168.1.110';
  port = "8080";
  getUrl: string;
  addUrl: string;
  delUrl: string;

  constructor(private http: HttpClient) { 
    // Golang API
    let baseUrl = "http://" + this.ip + ":" + this.port;

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
				},
				method: 'GET', 
				mode: '*' 
			};  
  
    return this.http.get<Song[]>(this.getUrl, options)
	  .pipe(
        catchError(this.handleError)
	  );
  }  
 
  addSong(newSong: NewSong) {
	  return this.http.post<NewSong>(this.addUrl, JSON.stringify(newSong))    
	    .pipe(catchError(this.handleError));
  } 
 
  deleteSong(id: number) {
    let options = {
          headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*',
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
