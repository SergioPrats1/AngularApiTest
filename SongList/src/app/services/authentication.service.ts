import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { User } from '../models';
import { environment } from '../../environments/environment';

@Injectable({ providedIn: 'root' })
export class AuthenticationService {
    private currentUserSubject: BehaviorSubject<User>;
    public currentUser: Observable<User>;

    TOKEN_KEY = 'token';

    constructor(private http: HttpClient) {
        this.currentUserSubject = new BehaviorSubject<User>(JSON.parse(localStorage.getItem('currentUser')));
        this.currentUser = this.currentUserSubject.asObservable();
    }

    public get currentUserValue(): User {
        return this.currentUserSubject.value;
    }

    login(username, password) {

        const headers = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
        };
        
        return this.http.post<any>(`${environment.apiEndPoint}/users/authenticate`, { username, password }, headers)
            .pipe(map(user => {
                // store user details and jwt token in local storage to keep user logged in between page refreshes
                user.password = "";
                localStorage.setItem('currentUser', JSON.stringify(user));
                this.currentUserSubject.next(user);
                localStorage.setItem(this.TOKEN_KEY, user.token);
                console.log("TOKEN retrieved after logging: " + user.token)
                return user;
            }));
    }

    logout() {
        // remove user from local storage and set current user to null
        localStorage.removeItem('currentUser');
        this.currentUserSubject.next(null);
        //Remove the TOKEN key
        localStorage.removeItem(this.TOKEN_KEY);
    }

    get token() {
        return localStorage.getItem(this.TOKEN_KEY);
    }

    get isAuthenticated() {
        return !!localStorage.getItem(this.TOKEN_KEY);
    }

}