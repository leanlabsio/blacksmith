import {Component} from "angular2/core";
import {View} from "angular2/core";
import {EventEmitter} from "angular2/core";
import {Job} from "./dashboard";
import {ChangeDetectionStrategy} from "angular2/core";
import {Observable} from "rxjs/Observable";
import {Input} from "angular2/core";
import {OnInit} from "angular2/core";
import {Env} from "./dashboard";
import {FORM_DIRECTIVES} from "angular2/common";
import {Headers} from "angular2/http";
import {Inject} from "angular2/core";
import {Http} from "angular2/http";
import {Builder} from "./dashboard";
import {RouteParams} from "angular2/router";


@Component({
    selector: "job-form"
})
@View({
    templateUrl: "html/job.form.html",
    directives: [FORM_DIRECTIVES],
})
export class JobForm implements OnInit
{
    job: Job;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) private params: RouteParams) {
        let hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        this.http.get("/api/jobs/"+params.get("repo"), {headers:hs})
            .map(res => Job.create(res.json()))
            .subscribe(job => this.job = job);
    }

    ngOnInit() {
        this.job = new Job({builder:new Builder({})});
    }

    addenv() {
        this.job.env.push(new Env({}));
    }

    save() {
        var hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        console.log(this.job);
        this.http.put("/api/jobs", JSON.stringify(this.job), {headers:hs})
            .map(res => res.json())
            .subscribe(val => console.log(val));
    }
}
