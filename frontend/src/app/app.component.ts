import { Component } from '@angular/core';
import { map, Observable, of, tap } from "rxjs";
import { Picture } from "./models";
import { GetPicturesGQL } from "../generated/graphql";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  loading = true;
  error: any;
  pictures$: Observable<Picture[]> = of([]);

  constructor(private query: GetPicturesGQL) {
  }

  ngOnInit() {
    this.pictures$ = this.query.watch().valueChanges.pipe(
      tap(res => this.loading = res.loading),
      tap(res => this.error = res.error),
      map(res => res.data.GetPictures),
      map(res => res as Picture[])
    );
  }
}
