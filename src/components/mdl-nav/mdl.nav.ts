import {
    Component,
    Directive,
    ContentChild
} from "@angular/core";


const template: string = <string>require('./mdl.nav.html');

@Component({
    selector: 'nav',
    template: template,
    host: {
        '[class.mdl-layout__header]': 'true'
    }
})
export class Navigation {
}

export const NAVIGATION_DIRECTIVES: any[] = [
    Navigation,
];
