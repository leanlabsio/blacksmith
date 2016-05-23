import {
    Component,
    Inject
} from "@angular/core";

import {
    Http,
    Headers
} from "@angular/http";

import {
    Router,
    RouteParams,
    ROUTER_DIRECTIVES
} from "@angular/router-deprecated";

import {Navigation} from "./../mdl-nav/mdl.nav";

import {
    Project,
    Repository,
    Trigger,
    Image,
    DockerExecutor,
    Env
} from "./../dashboard/dashboard";

const template: string = <string>require('./repo.browser.html');

@Component({
    selector: 'repo-browser',
    template: template,
    directives: [ROUTER_DIRECTIVES, Navigation],
})
export class RepoBrowser {

    jobs: Array<Project>;

    constructor(@Inject(Http) public http: Http, @Inject(Router) private router: Router, @Inject(RouteParams) private params: RouteParams) {
        var hs = new Headers();
        var qs = '?enabled=0';
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.get('/api/projects'+qs, {headers: hs}).map((res) => {var resp: Array<Job> = res.json(); return resp;}).subscribe(jobs => this.jobs = jobs);
    }

    enable(params: any, job: Project) {
        var hs = new Headers();
        job.trigger.active = true;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http
            .put("/api/projects/"+params.host+"/"+params.namespace+"/"+params.name, JSON.stringify(job), {headers: hs})
            .map((res) => {let resp:Project = res.json(); return resp;})
            .subscribe(res => this.router.navigate(['JobSettings', params]));
    }

    disable(params: any, job: Project) {
        var hs = new Headers();
        job.trigger.active = false;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.put("/api/projects/"+params.host+"/"+params.namespace+"/"+params.name, JSON.stringify(job), {headers: hs})
            .map(res => res.json())
            .map(data => <Project>data)
            .subscribe(val => job = val)
    }
}
