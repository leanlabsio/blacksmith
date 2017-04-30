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

const template: string = <string>require('./dashboard.html');

export interface Project {
    trigger: Trigger;
    executor: DockerExecutor;
    repository: Repository;
}

export interface Trigger {
    active: boolean;
}

export interface Repository {
    name: string;
    full_name: string;
    clone_url: string;
    description: string;
}

export interface DockerExecutor {
    image: Image;
    env: Array<Env>;
}

export class Image {
    name: string;
    tag: string;
}

export interface Env {
    name: string;
    value: string;
}

@Component({
    selector: 'dashboard',
    template: template,
    directives: [ROUTER_DIRECTIVES, Navigation],
})
export class Dashboard {

    jobs: Array<Project>;

    constructor(@Inject(Http) public http: Http, @Inject(Router) private router: Router, @Inject(RouteParams) private params: RouteParams) {
        var hs = new Headers();
        var qs = '?enabled=1';
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
