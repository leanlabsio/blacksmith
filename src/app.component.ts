import {
    Component,
    Inject
} from '@angular/core';

import {
    ROUTER_DIRECTIVES,
    RouteConfig,
    Router
} from '@angular/router-deprecated';

import {Home} from './components/home/home';
import {Dashboard} from "./components/dashboard/dashboard";
import {JobSettings} from "./components/job-settings/job.settings";
import {BuildList} from "./components/build-list/build.list.ts";
import {BuildLog} from "./components/build-log/build.log.ts";
import {RepoBrowser} from "./components/repo-browser/repo.browser";

@RouteConfig([
    {path: '/', component: Home, name: 'Home'},
    {path: '/jobs', component: Dashboard, name: 'Dashboard'},
    {path: '/jobs/create', component: RepoBrowser, name: 'DashboardCreate'},
    {path: '/log/:host/:namespace/:name', component: BuildList, name: 'BuildList'},
    {path: '/jobs/:host/:namespace/:name/settings', component: JobSettings, name: 'JobSettings'},
    {path: '/logs/:host/:namespace/:name/:commit/:timestamp', component: BuildLog, name: 'BuildLog'}
])
@Component({
    selector: 'app',
    template: <string>require('./app.component.html'),
    directives: [ROUTER_DIRECTIVES]
})
export class AppComponent{}
