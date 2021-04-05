import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CsCallerService {

  constructor(private CsProgram:HttpClient) {}

  getSongList(){
    return this.CsProgram.get('http://localhost:5000/api/songs')
  //TODO: Create this api endpoint in nginx and that endpoint calls our go program and then magic happens.
  }
}
