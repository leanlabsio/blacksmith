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


@Component({
    changeDetection: ChangeDetectionStrategy.OnPushObserve,
    selector: "job-form"
})
@View({
    templateUrl: "html/job.form.html",
    directives: [FORM_DIRECTIVES],
})
export class JobForm implements OnInit
{
    @Input() job: Job;

    constructor(@Inject(Http) private http: Http) {}

    ngOnInit() {
        this.job = new Job({});
    }

    addenv() {
        console.log(this.job);
        this.job.env.push(new Env({}));
    }

    save() {
        var hs = new Headers();
        hs.append("Authorization", "Bearer "+localStorage.getItem("jwt"));
        this.http.put("/jobs", JSON.stringify(this.job), {headers:hs})
        .map(res => res.json())
        .subscribe(val => console.log(val));
    }
}