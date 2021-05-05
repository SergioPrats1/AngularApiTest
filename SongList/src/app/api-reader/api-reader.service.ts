import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

export interface Song {
  id: number;
  title: string;
  artist: string;
	year: number;
}

@Injectable({
  providedIn: 'root'
})
export class ApiReaderService {

  constructor(private http: HttpClient) { }
  
  //configUrl = 'http://localhost:5000/api/songs';	  
  configUrl = 'assets/Songs.json';
  
  getSongs() {   
  
  let options = {
				headers: {
					'Content-Type': 'application/json',
				},
				method: 'GET', // GET, POST, PUT, DELETE
				mode: 'no-cors' // the most important option
			};  
  
    return this.http.get<Song[]>(this.configUrl, options);
  }  
 
  
}
