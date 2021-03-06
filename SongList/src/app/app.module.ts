import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
//import { HttpModule } from '@angular/http';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MySongListComponent } from './my-song-list/my-song-list.component';

@NgModule({
  declarations: [
    AppComponent,
    MySongListComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
	HttpClientModule
	//HttpModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
