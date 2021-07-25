import { Component, OnInit } from '@angular/core';



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {

  showSongList: Boolean;

  constructor(){
	  this.showSongList = false;
  }

  ngOnInit(){
  }

  loadSongList() {
    this.showSongList = true;
  }


  listClosed(isClosed) {
    console.log("listClosed() was called");
    this.showSongList = false;
  }

}
