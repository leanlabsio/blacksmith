import {Component} from '@angular/core';
import {Login} from './../login/login';

@Component({
  selector: 'home',
  directives: [Login],
  template: <string>require('./home.html')
})
export class Home{
  private ghclientid: string;

  constructor() {
    this.ghclientid = window.bsconfig.github.oauth.clientid;
  }
}
