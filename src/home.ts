import {Component} from 'angular2/core';
import {Login} from './login';

@Component({
  directives: [Login],
  template: '<login [ghclient]="ghclientid"></login>'
})
export class Home{
  private ghclientid: string;

  constructor() {
    this.ghclientid = window.bsconfig.github.oauth.clientid;
  }
}
