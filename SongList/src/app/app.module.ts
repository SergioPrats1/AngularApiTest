import { RouterModule , Routes } from '@angular/router';

import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { MySongListComponent } from './my-song-list/my-song-list.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatTableModule} from '@angular/material/table';
import { AddSongComponent } from './add-song/add-song.component'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Error404Component } from './error404/error404.component';
import { UserLoginComponent } from './user-login/user-login.component';
import { HomePageComponent } from './home-page/home-page.component';
import { UserRegisterComponent } from './user-register/user-register.component';


const rutasApp:Routes = [
  { path:'song-list' , component: MySongListComponent },
  { path:'add-song' , component: AddSongComponent },
  { path:'user-login' , component: UserLoginComponent },
  { path:'details' , redirectTo: 'song-list' },
  { path:'404' , component: Error404Component },
  { path:'' , component: HomePageComponent , pathMatch: 'full' },  
  { path:'register' , component: UserRegisterComponent },
  { path:'**' , redirectTo: '404' }
]

@NgModule({
  declarations: [
    AppComponent,
    MySongListComponent,
    AddSongComponent,
    Error404Component,
    UserLoginComponent,
    HomePageComponent,
    UserRegisterComponent
  ],
  imports: [
    RouterModule.forRoot(rutasApp),    
    BrowserModule,
	  HttpClientModule,
	  BrowserAnimationsModule,
	  MatTableModule,
	  ReactiveFormsModule
  ],
  providers: [HttpClientModule],
  bootstrap: [AppComponent]
})
export class AppModule { }
