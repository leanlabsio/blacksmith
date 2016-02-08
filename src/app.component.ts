import {Component, View} from 'angular2/core';
import {Home} from './home';
import {ROUTER_DIRECTIVES, RouteConfig, Router} from 'angular2/router';
import {Dashboard} from "./dashboard";
import {JobPage} from "./job";

@RouteConfig([
    {path: '/', component: Home, name: 'Home'},
    {path: '/dashboard', component: Dashboard, name: 'Dashboard'},
    {path: '/job/:repo', component: JobPage, name: 'Job'}
])
@Component({
    selector: 'app'
})
@View({
    template: '<router-outlet></router-outlet>',
    directives: [ ROUTER_DIRECTIVES ]
})
export class AppComponent{}
