import {Component} from '@angular/core';
import {Apollo, gql} from "apollo-angular";
import {map, Observable, of, tap} from "rxjs";
import {Picture} from "./models";

const GetPictures = gql`
  query GetPictures {
    GetPictures {
      id, path
    }
  }
`

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  loading = true;
  error: any;
  pictures$: Observable<Picture[]> = of([]);

  constructor(private apollo: Apollo) {
  }

  ngOnInit() {
    this.pictures$ = this.apollo.watchQuery<any>({query: GetPictures}).valueChanges.pipe(
      tap(res => this.loading = res.loading),
      tap(res => this.error = res.error),
      map(res => res.data.GetPictures),
      map(res => res as Picture[])
    );
  }
}
