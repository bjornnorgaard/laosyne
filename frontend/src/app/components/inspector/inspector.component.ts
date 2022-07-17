import { Component, OnInit } from '@angular/core';
import { InspectorSearchGQL, InspectorSearchQuery } from "../../../generated/graphql";
import { map, Observable, tap } from "rxjs";

@Component({
  selector: 'app-inspector',
  templateUrl: './inspector.component.html',
  styleUrls: ['./inspector.component.scss']
})
export class InspectorComponent implements OnInit {
  public loading: boolean = true;
  public error: any = null;
  public result$: Observable<InspectorSearchQuery> | undefined;

  constructor(private query: InspectorSearchGQL) {
  }

  ngOnInit(): void {
    this.update();
  }

  private update(): void {
    this.result$ = this.query.watch({input: {take: 100}}).valueChanges.pipe(
      tap(res => this.loading = res.loading),
      tap(res => this.error = res.error),
      map(res => res.data)
    );
  }
}
