import {Component} from "angular2/core";
import {JobForm} from "./../job-form/job.form.ts";
import {Inject} from "angular2/core";
import {ElementRef} from "angular2/core";
import {OnInit} from "angular2/core";
import {AfterViewInit} from "angular2/core";
import {Navigation} from "./../navigation/navigation";

const template: string = <string>require('./job.settings.html');

@Component({
    template: template,
    directives: [JobForm, Navigation]
})
export class JobSettings implements OnInit {

    constructor(@Inject(ElementRef) private element: ElementRef) {}

    ngOnInit() {
//        window.componentHandler.upgradeElements(this.element.nativeElement);
    }
}
