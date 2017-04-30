import {
    Component,
    Inject,
    OnInit,
    ElementRef
} from "@angular/core";

import {JobForm} from "./../job-form/job.form.ts";
import {NAVIGATION_DIRECTIVES} from "./../mdl-nav/mdl.nav";

const template: string = <string>require('./job.settings.html');

@Component({
    template: template,
    directives: [JobForm, NAVIGATION_DIRECTIVES]
})
export class JobSettings implements OnInit {

    constructor(@Inject(ElementRef) private element: ElementRef) {}

    ngOnInit() {
//        window.componentHandler.upgradeElements(this.element.nativeElement);
    }
}
