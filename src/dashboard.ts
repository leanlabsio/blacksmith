import {Component} from "angular2/core";
import {View} from "angular2/core";
import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import {Inject} from "angular2/core";
import {RequestOptions} from "https";
import {Headers} from "angular2/http";
import {Router} from "angular2/router";
import {ROUTER_DIRECTIVES} from "angular2/router";

export class Job {
    repository: string;
    enabled: boolean;
    env: Array<Env>;
    name: string;

    static create(data) {
        return new Job(data);
    }

    constructor(data) {
        this.name = data.name;
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
    template: `
    <div *ngFor="#job of jobs" class="row align-center">
        <div class="columns medium-8">
            <a [routerLink]="['BuildList', {repo: job.repository}]">
                {{job.name}}
            </a>
        </div>
        <div class="columns medium-4">
            <button *ngIf="job.enabled == false" class="button success" (click)="enable(job)">
                Enable build
            </button>
            <button *ngIf="job.enabled == true" class="button alert" (click)="disable(job)">
                Disable build
            </button>
        </div>
    </div>
    `,
    directives: [ROUTER_DIRECTIVES],
})
export class Dashboard {

    jobs: Array<Job>;

    constructor(@Inject(Http) public http: Http, @Inject(Router) private router: Router) {
        var hs = new Headers();
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.get('/api/jobs', {headers: hs}).map(res => res.json()).subscribe(jobs => this.jobs = jobs);
    }

    enable(job: Job) {
        var hs = new Headers();
        job.enabled = true;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http
            .put("/api/jobs",JSON.stringify(job), {headers: hs})
            .map(res => res.json())
            .subscribe(res => this.router.navigate(['Job', {repo: res.repository}]));
    }

    disable(job: Job) {
        var hs = new Headers();
        job.enabled = false;
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        this.http.put("/api/jobs", JSON.stringify(job), {headers: hs})
        .map(res => res.json())
        .subscribe(val => job = Job.create(val))
    }
}