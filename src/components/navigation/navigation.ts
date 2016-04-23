import {
    Component,
    Directive,
    ContentChild
} from "angular2/core";


const template: string = <string>require('./navigation.html');

@Component({
    selector: 'nav',
    template: template
})
export class Navigation {
}

export const NAVIGATION_DIRECTIVES: any[] = [
    Navigation,
];
