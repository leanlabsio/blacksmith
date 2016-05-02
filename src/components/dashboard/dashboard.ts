import {
    Component,
    Inject
} from "angular2/core";

import {
    Http,
    Headers
} from "angular2/http";

import {
    Router,
    ROUTER_DIRECTIVES
} from "angular2/router";

import {Navigation} from "./../mdl-nav/mdl.nav";

const template: string = <string>require('./dashboard.html');

export interface Project {
    enabled: boolean;
    env: Array<Env>;
    builder: Builder;
    repository: Repository;
}

export interface Repository {
    name: string;
    full_name: string;
    clone_url: string;
    description: string;
}

export class Builder {
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

    constructor(@Inject(Http) public http: Http, @Inject(Router) private router: Router) {
        var hs = new Headers();
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.get('/api/jobs', {headers: hs}).map((res) => {var resp: Array<Job> = res.json(); return resp;}).subscribe(jobs => this.jobs = jobs);
    }

    enable(job: Job) {
        var hs = new Headers();
        job.enabled = true;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http
            .put("/api/jobs", JSON.stringify(job), {headers: hs})
            .map((res) => {let resp:Project = res.json(); return resp;})
            .subscribe(res => this.router.navigate(['Job', {repo: res.repository.clone_url}]));
    }

    disable(job: Job) {
        var hs = new Headers();
        job.enabled = false;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.put("/api/jobs", JSON.stringify(job), {headers: hs})
            .map(res => res.json())
            .map(data => <Project>data)
            .subscribe(val => job = val)
    }
}
