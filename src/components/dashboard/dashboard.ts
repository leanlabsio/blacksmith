import {Component} from "angular2/core";
import {Http} from "angular2/http";
import {Inject} from "angular2/core";
import {Headers} from "angular2/http";
import {Router} from "angular2/router";
import {ROUTER_DIRECTIVES} from "angular2/router";
import {Navigation} from "./../navigation/navigation";

const template: string = <string>require('./dashboard.html');

export interface Job {
    repository: string;
    enabled: boolean;
    env: Array<Env>;
    name: string;
    full_name: string;
    clone_url: string;
    builder: Builder;
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

    jobs: Array<Job>;

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
            .map((res) => {let resp:Job = res.json(); return resp;})
            .subscribe(res => this.router.navigate(['Job', {repo: res.clone_url}]));
    }

    disable(job: Job) {
        var hs = new Headers();
        job.enabled = false;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.put("/api/jobs", JSON.stringify(job), {headers: hs})
            .map(res => res.json())
            .map(data => <Job>data)
            .subscribe(val => job = val)
    }
}
