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
  
  // C# API (legacy)
  //configUrl = 'http://localhost:5000/api/songs';	  
  
  // Golang API
  configUrl = 'http://localhost:8080/getSongs';
  //configUrl = 'assets/Songs.json';
  
  
  getSongs() {   
  
  let options = {
				headers: {
					'Content-Type': 'application/json',
					'Access-Control-Allow-Origin': '*',
				},
				method: 'GET', // GET, POST, PUT, DELETE
				mode: '*' // the most important option
			};  
  
    return this.http.get<Song[]>(this.configUrl, options);
  }  
 
  
}
