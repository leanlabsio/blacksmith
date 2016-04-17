import {Component} from 'angular2/core';
import {Login} from './../login/login';

@Component({
  directives: [Login],
  template: <string>require('./home.html')
})
export class Home{
  private ghclientid: string;

  constructor() {
    this.ghclientid = window.bsconfig.github.oauth.clientid;
  }
}
