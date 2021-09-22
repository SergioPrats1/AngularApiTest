import { Component, OnInit } from '@angular/core';

export class User {
  UserName: string;
  Password: string;
  email: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {

  constructor(){
  }

  ngOnInit(){
  }

}
