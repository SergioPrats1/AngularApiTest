import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { Router } from '@angular/router';
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
  displayedColumns: string[] = ['title', 'artist', 'year', 'delete'];
  @Output() onClose: EventEmitter<boolean>;

  constructor(private apiReaderService: ApiReaderService, private readonly router: Router){
	this.loading = false;
    this.errorHttp = false;
    this.onClose = new EventEmitter();
  }

  ngOnInit() {
    this.loading = true;
	  this.showSongs();
	  this.loading = false;
	  this.songs = [];
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

  closeList(){
    console.log("closeList() was called");

    // I am no longer using events.
    //this.onClose.emit(true);
    this.router.navigate(['']);
  }
  
  loadAddSong() {
    this.router.navigate(['add-song']);
  }
  
  AddSongClosed(isClosed) {
    console.log("listClosed() was called");
    this.showSongs();
  }

  DeleteSong(Id: number) {

    let proceed =confirm('Are you sure you wish to delete?')

    if (proceed) {
      this.apiReaderService.deleteSong(Id)
	      .subscribe(() => {  this.showSongs();  });
    }
  }

}
