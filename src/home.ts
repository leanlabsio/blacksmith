import {Component, View} from 'angular2/core';
import {Login} from './login';

@Component({
})
@View({
    directives: [Login],
    template: '<login></login>'
})
export class Home{}
