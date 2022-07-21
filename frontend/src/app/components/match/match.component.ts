import { Component, HostListener, OnInit } from '@angular/core';
import { CreateMatchGQL, CreateMatchQuery, ReportMatchWinnerGQL } from "../../../generated/graphql";
import { filter, map, Observable, of, tap } from "rxjs";

export enum KEY_CODE {
  RIGHT_ARROW = 'ArrowRight',
  LEFT_ARROW = 'ArrowLeft'
}

@Component({
  selector: 'app-match',
  templateUrl: './match.component.html',
  styleUrls: ['./match.component.scss']
})
export class MatchComponent implements OnInit {
  public result$: Observable<CreateMatchQuery> = of({} as CreateMatchQuery);

  // Only used by key listener func.
  private playerTwoId: number = 0;
  private playerOneId: number = 0;

  constructor(private match: CreateMatchGQL,
              private report: ReportMatchWinnerGQL) {
  }

  ngOnInit(): void {
    this.createMatch();
  }

  private createMatch() {
    this.result$ = this.match.watch().valueChanges.pipe(
      filter(res => !!res.data.Match.playerOne.id),
      map(res => res.data),
      tap(data => this.playerOneId = data.Match.playerOne.id),
      tap(data => this.playerTwoId = data.Match.playerTwo.id),
    );
  }

  public reportResult(winnerId: number, loserId: number): void {
    this.report.mutate({winnerId: winnerId, loserId: loserId}).subscribe();
    this.match.watch().refetch();
  }

  @HostListener('window:keyup', ['$event'])
  keyEvent(event: KeyboardEvent): void {
    if (event.key === KEY_CODE.RIGHT_ARROW || event.key === 'd') {
      this.reportResult(this.playerTwoId, this.playerOneId);
    } else if (event.key === KEY_CODE.LEFT_ARROW || event.key === 'a') {
      this.reportResult(this.playerOneId, this.playerTwoId);
    }
  }

}
