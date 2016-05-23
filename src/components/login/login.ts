import {
    Component,
    Inject,
    Input,
} from '@angular/core';
import {Http} from '@angular/http';
import 'rxjs/Rx';
import {Router} from "@angular/router-deprecated";
import * as uri from 'urijs';

const template: string = <string>require('./login.html');

@Component({
    selector: 'login',
    template: template
})
export class Login{
    @Input() ghclient: string;

    constructor(@Inject(Http) private http: Http, @Inject(Router) public router: Router) {}

    authenticate(provider: string) {
        let scopes = ["user:email", "write:repo_hook", "repo", "admin:repo_hook"];
        let scope = scopes.join(",");
        let popup = window.open('https://github.com/login/oauth/authorize?client_id='+this.ghclient+"&scope="+scope);
        let redirectUri = window.location.origin + '/oauth';
        let http = this.http;
        let router = this.router;

        let handle = setInterval(function() {
            let loc = popup.location.href;
            if (loc.indexOf(redirectUri) != -1) {
                let url = new uri(popup.location.href);
                let search = url.search(true);
                let j = JSON.stringify({code: search.code});
                http.post('/api/auth/github', j)
                    .map(res => res.json())
                    .subscribe(value => {localStorage.setItem("jwt", value.token); router.navigate(['Dashboard']) });

                popup.close();
                clearInterval(handle);
            }
        }, 200);

    }
}
