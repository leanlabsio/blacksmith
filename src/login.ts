import {Component, View} from 'angular2/core';
import {parser} from 'angular2/src/router/url_parser';
import {Http} from 'angular2/http';
import 'rxjs/Rx';
import {Inject} from "angular2/core";
import {Router} from "angular2/router";
import {Input} from "angular2/core";

@Component({
    selector: 'login'
})
@View({
    templateUrl: 'html/login.html'
})
export class Login{
    @Input() ghclient: string;

    constructor(@Inject(Http) private http: Http, @Inject(Router) public router: Router) {}

    authenticate(provider: string) {
        let popup = window.open('https://github.com/login/oauth/authorize?client_id='+this.ghclient);
        let redirectUri = window.location.origin + '/';
        let http = this.http;
        let router = this.router;

        let handle = setInterval(function() {
            let uri = popup.location.protocol + '//' + popup.location.host + popup.location.pathname;
            if (redirectUri === uri) {
                let url = parser.parse(popup.location.search);
                let j = JSON.stringify({code: url.params.code});
                http.post('/auth/github', j)
                    .map(res => res.json())
                    .subscribe(value => {localStorage.setItem("jwt", value.token); router.navigate(['Dashboard', {}]) });

                popup.close();
                clearInterval(handle);
            }
        }, 200);

    }
}
