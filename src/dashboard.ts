import {Component} from "angular2/core";
import {View} from "angular2/core";
import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import {Inject} from "angular2/core";
import {RequestOptions} from "https";
import {Headers} from "angular2/http";
import {Router} from "angular2/router";

export class Job {
    repository: string;
    enabled: boolean;
    env: Array<Env>;

    static create(data) {
        return new Job(data);
    }

    constructor(data) {
        this.repository = data.repository;
        this.enabled = data.enabled;
        this.env = [];
    }
}

export class Env {
    name: string;
    value: string;

    constructor(data) {
        this.name = data.name;
        this.value = data.value;
    }
}

@Component({})
@View({
    templateUrl: "html/dashboard.html",
})
export class Dashboard {

    jobs: Array<Job>;

    constructor(@Inject(Http) public http: Http, @Inject(Router) private router: Router) {
        var hs = new Headers();
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.get('/jobs', {headers: hs}).map(res => res.json()).subscribe(jobs => this.jobs = jobs);
    }

    enable(job: Job) {
        var hs = new Headers();
        job.enabled = true;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http
            .put("/jobs",JSON.stringify(job), {headers: hs})
            .map(res => res.json())
            .subscribe(res => this.router.navigate(['Job', {repo: res.repository}]));
    }
}