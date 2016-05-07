import {Component} from 'angular2/core';
import {Home} from './components/home/home';
import {ROUTER_DIRECTIVES, RouteConfig, Router} from 'angular2/router';
import {Dashboard} from "./components/dashboard/dashboard";
import {JobSettings} from "./components/job-settings/job.settings";
import {BuildList} from "./components/build-list/build.list.ts";
import {BuildLog} from "./components/build-log/build.log.ts";

@RouteConfig([
    {path: '/', component: Home, name: 'Home'},
    {path: '/jobs', component: Dashboard, name: 'Dashboard'},
    {path: '/jobs/:host/:namespace/:name', component: BuildList, name: 'BuildList'},
    {path: '/jobs/:host/:namespace/:name/settings', component: JobSettings, name: 'Job'},
    {path: '/jobs/:repo/:commit', component: BuildLog, name: 'BuildLog'}
])
@Component({
    selector: 'app',
    template: <string>require('./app.component.html'),
    directives: [ROUTER_DIRECTIVES]
})
export class AppComponent{}
