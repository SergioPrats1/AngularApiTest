import { Component, OnInit } from '@angular/core';
import { Router} from '@angular/router';
import {User } from '../app.component';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  title: string;
  currentUser: User;

  constructor(private readonly router: Router){
	  this.title="SongList";
    this.currentUser = new User();
    this.currentUser.UserName = 'Anonymous';
  }

  ngOnInit(): void {
  }

  loadSongList() {
    this.router.navigate(['song-list']);
  }

  listClosed(isClosed) {
    console.log("listClosed() was called");
  }

}
