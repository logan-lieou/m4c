import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  readonly ROOT_URL = "http://localhost:8080/api/products";

  post: any;

  constructor ( private http: HttpClient ) {  }

  getPosts() {
    this.http.get(this.ROOT_URL, {responseType: 'text'})
      .subscribe(response => this.post = response);
  }
}
