import {Component} from "angular2/core";
import {View} from "angular2/core";
import {Inject} from "angular2/core";
import {RouteParams} from "angular2/router";
import {Job} from "./dashboard";
import {Http} from "angular2/http";
import {Headers} from "angular2/http";
import {JobForm} from "./job.form";
import {ChangeDetectionStrategy} from "angular2/core";

@Component({
    changeDetection: ChangeDetectionStrategy.OnPushObserve
})
@View({
    template: `
    <job-form [job]="job"></job-form>
    `,
    directives: [JobForm]
})
export class JobPage {
    private job: Job;
    constructor(@Inject(RouteParams) private params: RouteParams, @Inject(Http) private http: Http) {
        var hs = new Headers();
        hs.append("Authorization", "Bearer " + localStorage.getItem("jwt"));
        http.get("/jobs/"+params.get("repo"), {headers: hs}).map(res => res.json()).subscribe(val => this.job = val);
    }
}