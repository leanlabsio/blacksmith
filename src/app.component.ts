import {Component} from 'angular2/core';
import {Home} from './components/home/home';
import {ROUTER_DIRECTIVES, RouteConfig, Router} from 'angular2/router';
import {Dashboard} from "./components/dashboard/dashboard";
import {JobPage} from "./components/job/job";
import {BuildList} from "./components/build-list/build.list.ts";
import {BuildLog} from "./components/build-log/build.log.ts";

@RouteConfig([
    {path: '/', component: Home, name: 'Home'},
    {path: '/jobs', component: Dashboard, name: 'Dashboard'},
    {path: '/jobs/:repo', component: BuildList, name: 'BuildList'},
    {path: '/jobs/:repo/settings', component: JobPage, name: 'Job'},
    {path: '/jobs/:repo/:commit', component: BuildLog, name: 'BuildLog'}
])
@Component({
    selector: 'app',
    template: '<router-outlet></router-outlet>',
    directives: [ ROUTER_DIRECTIVES ]
})
export class AppComponent{}
