import {Component} from 'angular2/core';
import {Login} from './../login/login';

@Component({
  directives: [Login],
  template: require('./home.html')
})
export class Home{
  private ghclientid: string;

  constructor() {
    this.ghclientid = window.bsconfig.github.oauth.clientid;
  }
}
