import { Component, OnInit } from '@angular/core';
import { Song, ApiReaderService } from '../api-reader/api-reader.service';

@Component({
  selector: 'app-my-song-list',
  templateUrl: './my-song-list.component.html',
  styleUrls: ['./my-song-list.component.css']
})
export class MySongListComponent implements OnInit {
  title = 'SongList';
  songs:Array<Song>;
  loading:Boolean;
  errorHttp: Boolean;
  error: any;
  song:Song;
  
  constructor(private apiReaderService: ApiReaderService){
	this.loading=false;
    this.errorHttp=false;
  }
  
  ngOnInit() {
    this.loading = true;
	this.showSongs();
	this.loading = false;
  }

  showSongs() { 
	this.apiReaderService.getSongs()
		.subscribe((data : Song[]) => { this.songs = data; console.log(data); 
											this.loading = false; },
					error => this.error = error	  );		
  }


  showSong(_song: Song){
    alert( ` ${_song.title} is from ${_song.artist} ` )
  }

  errorHttpColor(){
	  if( this.errorHttp == true) {
		  return "#ff0000";
	  }
	  else {
		  return "#0000ff";
	  }      		  
  }

}
