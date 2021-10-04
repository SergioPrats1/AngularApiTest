import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { first } from 'rxjs/operators';

import { AuthenticationService } from '../services';
import { User } from '../models';


@Component({
  selector: 'app-user-login',
  templateUrl: './user-login.component.html',
  styleUrls: ['./user-login.component.css']
})

export class UserLoginComponent implements OnInit {
  loginForm: FormGroup;
  loading = false;
  submitted = false;
  returnUrl: string;

  user: User;

  constructor(
      private formBuilder: FormBuilder,
      private route: ActivatedRoute,
      private router: Router,
      private authenticationService: AuthenticationService
  ) {
      // redirect to home if already logged in
      if (this.authenticationService.currentUserValue) {
        this.user =  <User>(JSON.parse(localStorage.getItem('currentUser')));
        console.log("Hi " + this.user.firstname)
      }
  }

  ngOnInit() {
      this.loginForm = this.formBuilder.group({
          username: ['', Validators.required],
          password: ['', Validators.required]
      });

      // get return url from route parameters or default to '/'
      this.returnUrl = 'song-list';
  }

  // convenience getter for easy access to form fields
  get f() { return this.loginForm.controls; }

  onSubmit() {
      this.submitted = true;

      // stop here if form is invalid
      if (this.loginForm.invalid) {
          return;
      }

      this.loading = true;
      this.authenticationService.login(this.f.username.value, this.f.password.value)
          .pipe(first())
          .subscribe(
              data => {
                this.user =  <User>(JSON.parse(localStorage.getItem('currentUser')));
                console.log(this.user.email);
                alert ("Hello " + this.user.firstname + " " + this.user.lastname);
                  this.router.navigate([this.returnUrl]);
              },
              error => {
                  console.log("there has been an error: " + error);
                  this.loading = false;
              });
  }
}
