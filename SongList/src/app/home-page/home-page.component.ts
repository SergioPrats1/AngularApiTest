import { Component, OnInit } from '@angular/core';
import { Router} from '@angular/router';
import { AuthenticationService } from '../services';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  title: string;

  UserIsLogged: boolean;


  constructor(private readonly router: Router,
              private authenticationService: AuthenticationService){
	  this.title="SongList";
  }

  ngOnInit(): void {
    this.UserIsLogged = this.authenticationService.isAuthenticated;
  }

  loadSongList() {
    this.router.navigate(['song-list']);
  }

  /*listClosed(isClosed) {
    console.log("listClosed() was called");
  }*/

  login() {
    this.router.navigate(['user-login']);
  }  

  logout() {
    this.authenticationService.logout();
    this.UserIsLogged = false;
  }  

}
