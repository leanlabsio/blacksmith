import {Component} from "angular2/core";
import {EventEmitter} from "angular2/core";
import {Job} from "./../dashboard/dashboard";
import {ChangeDetectionStrategy} from "angular2/core";
import {Observable} from "rxjs/Observable";
import {Input} from "angular2/core";
import {OnInit} from "angular2/core";
import {Env} from "./../dashboard/dashboard";
import {FORM_DIRECTIVES} from "angular2/common";
import {Headers} from "angular2/http";
import {Inject} from "angular2/core";
import {Http} from "angular2/http";
import {Builder} from "./../dashboard/dashboard";
import {RouteParams} from "angular2/router";
import {MdInput} from "./../mdl-textfield/mdl.textfield";

@Component({
    selector: "job-form",
    template: <string>require('./job.form.html'),
    directives: [FORM_DIRECTIVES, MdInput],
})
export class JobForm implements OnInit
{
    job: Job;

    constructor(@Inject(Http) private http: Http, @Inject(RouteParams) private params: RouteParams) {
        let hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        this.http.get("/api/jobs/"+params.get("repo"), {headers:hs})
            .map(res => <Job>res.json())
            .subscribe(job => this.job = job);
    }

    ngOnInit() {
        let builder: Builder = {};
        let env: Env[] = [];
        this.job = <Job>({builder: builder, env: env});
    }

    addenv() {
        if (!this.job.env || !this.job.env.length) {
            let env: Env[] = [];
            this.job.env = env;
        }
        this.job.env.push(<Env>{});
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
