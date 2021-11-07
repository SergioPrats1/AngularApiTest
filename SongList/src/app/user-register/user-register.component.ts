import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators , ValidationErrors, ValidatorFn, AbstractControl} from '@angular/forms';
import { first } from 'rxjs/operators';
import { AuthenticationService, UserService } from '../services';
import { User } from '../models';

@Component({
  selector: 'app-user-register',
  templateUrl: './user-register.component.html',
  styleUrls: ['./user-register.component.css']
})
export class UserRegisterComponent implements OnInit {
  registerForm: FormGroup;
  loading = false;
  submitted = false;
  returnUrl: string;
  passwordsMatch: boolean;

  constructor(
      private formBuilder: FormBuilder,
      private router: Router,
      private authenticationService: AuthenticationService,
      private userService: UserService
  ) {
      // redirect to home if already logged in
      if (this.authenticationService.currentUserValue) {
          this.router.navigate(['/']);
      }
  }

  ngOnInit() {
      this.registerForm = this.formBuilder.group({
          userName: ['', Validators.required],
          password: ['', [Validators.required, Validators.minLength(6)]], 
          confirmPassword: [''],
          firstName: ['', Validators.required],
          lastName: ['', Validators.required],
          email: ['', Validators.required]
      },
      {validator: checkPasswords} 
      );
      this.returnUrl = 'user-login';
      this.passwordsMatch = true;
      this.loading = false
  }

  // convenience getter for easy access to form fields
  get f() { return this.registerForm.controls; }

  get f2() { return this.registerForm; }

  onSubmit() {
    this.submitted = true;
      var user = new User();

      // stop here if form is invalid
      if (this.registerForm.invalid) {
        this.getFormValidationErrors();
        console.log("invalid form");
        return;
      }

      this.passwordsMatch = true;
      user.username = this.registerForm.controls.userName.value;
      user.password = this.registerForm.controls.password.value;
      user.firstname = this.registerForm.controls.firstName.value;
      user.lastname = this.registerForm.controls.lastName.value;
      user.email = this.registerForm.controls.email.value;

      this.loading = true;
      this.userService.register(user)
          .pipe(first())
          .subscribe(
              data => {
                  this.router.navigate([this.returnUrl]);
              },
              error => {
                  this.loading = false;
              });
    }

    getFormValidationErrors() {
        Object.keys(this.registerForm.controls).forEach(key => {
      
            const controlErrors: ValidationErrors = this.registerForm.get(key).errors;
            if (controlErrors != null) {
                  Object.keys(controlErrors).forEach(keyError => {
                    console.log('Key control: ' + key + ', keyError: ' + keyError + ', err value: ', controlErrors[keyError]);
                });
            }
        });

        if (this.registerForm.errors?.['passMismatch'])
        {
            console.log("pass mismatch!")
        }
    }
}

export const checkPasswords: ValidatorFn = (control: AbstractControl): ValidationErrors | null => {
    const pass = control.get('password');
    const pass2 = control.get('confirmPassword');
    return pass && pass2 && pass.value != pass2.value ? { passMismatch: true } : null;
};
