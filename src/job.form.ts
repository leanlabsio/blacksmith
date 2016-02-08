import {Component} from "angular2/core";
import {View} from "angular2/core";
import {EventEmitter} from "angular2/core";
import {Job} from "./dashboard";
import {ChangeDetectionStrategy} from "angular2/core";
import {Observable} from "rxjs/Observable";

@Component({
    changeDetection: ChangeDetectionStrategy.OnPush,
    selector: "job-form",
    inputs: ["job"],
    outputs: ["jobchange"]
})
@View({
    templateUrl: "html/job.form.html"
})
export class JobForm {
    public jobchange: EventEmitter = new EventEmitter();
    public job: Observable<Job>;

    constructor() {
    }

    do() {
        console.log(this.job);

    }
}