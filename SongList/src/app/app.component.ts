import { Component, OnInit } from '@angular/core';
//import {Observable,of, from } from 'rxjs';
//import { Http, Response } from '@angular/http';
//import { HttpClient, HttpResponse } from '@angular/common/http';
import { Song, ApiReaderService } from './api-reader.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  title = 'SongList';
  songs:Array<Song>;
  loading:Boolean;
  errorHttp: Boolean;
  dummyText: string;
  error: any;
  song:Song

  constructor(private apiReaderService: ApiReaderService){
    /*this.songs = [
      {id:1,title:'ConstructorDefault_1', artist:'SergioP',year:2021},
      {id:2,title:'ConstructorDefault_2', artist:'SergioP',year:2021}
		];*/
	this.loading=false;
    this.errorHttp=false;
    this.dummyText='abc';
	this.song = {id: 0, title: '', artist: '', year: -1};
  }

  ngOnInit() {
    this.loading = true;
	// this.songs = this.showSongs();
	this.showSongs();
	this.loading = false;
  }

  showSongs() { 
	  
    /*this.http.request('http://localhost:5000/api/songs').subscribe(
      (respuesta: Response) => { this.dummyText = respuesta.json(); this.loading = false },
      (respuesta: Response) => { this.errorHttp = true } )*/
      console.log('zzz')
	  /*this.apiReaderService.getSongs()
		.subscribe((data : Song) => { this.song = { ...data }; console.log(data); 
											console.log(this.song);
											this.loading = false; 
											console.log(this.song.title);},
					error => this.error = error	  );*/
	this.apiReaderService.getSongs()
		.subscribe((data : Song[]) => { this.songs = data; console.log(data); 
											this.loading = false; },
					error => this.error = error	  );		
  }


  showSong(_song: Song){
    alert( ` ${_song.title} is from ${_song.artist} ` )
  }

}
