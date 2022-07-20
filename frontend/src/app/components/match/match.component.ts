import { Component, OnInit } from '@angular/core';
import { CreateMatchGQL, CreateMatchQuery, ReportMatchWinnerGQL } from "../../../generated/graphql";
import { map, Observable, of } from "rxjs";

@Component({
  selector: 'app-match',
  templateUrl: './match.component.html',
  styleUrls: ['./match.component.scss']
})
export class MatchComponent implements OnInit {
  public result$: Observable<CreateMatchQuery> = of({} as CreateMatchQuery);

  constructor(private match: CreateMatchGQL,
              private report: ReportMatchWinnerGQL) {
  }

  ngOnInit(): void {
    this.createMatch();
  }

  private createMatch() {
    this.result$ = this.match.watch().valueChanges.pipe(map(res => res.data));
  }

  public reportResult(winnerId: number, loserId: number): void {
    this.report.mutate({winnerId: winnerId, loserId: loserId}).subscribe();
    this.match.watch().refetch();
  }

}
