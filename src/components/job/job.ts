import {Component} from "angular2/core";
import {Inject} from "angular2/core";
import {RouteParams} from "angular2/router";
import {Http} from "angular2/http";
import {JobForm} from "./../job-form/job.form.ts";
import {CORE_DIRECTIVES} from "angular2/common";

@Component({
    template: `
    <job-form></job-form>
    `,
    directives: [JobForm, CORE_DIRECTIVES]
})
export class JobPage {
    public job:any;
    constructor(@Inject(RouteParams) private params: RouteParams, @Inject(Http) private http: Http) {
    }
}
