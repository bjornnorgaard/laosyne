import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { environment } from "../../../environments/environment";
import { PictureDetailsGQL, PictureDetailsQuery } from "../../../generated/graphql";
import { map, Observable, tap } from "rxjs";

@Component({
  selector: 'image-details',
  templateUrl: './image-details.component.html',
  styleUrls: ['./image-details.component.scss']
})
export class ImageDetails implements OnChanges {
  @Input() public id: number = 4;
  @Input() public full: boolean = false;

  public api: string = environment.api;
  public picture: Observable<PictureDetailsQuery> | undefined;
  public loading: boolean = true;
  public error: any = null;

  constructor(private query: PictureDetailsGQL) {
    this.update();
  }

  private update() {
    this.picture = this.query.watch({id: this.id}).valueChanges.pipe(
      tap(res => this.loading = res.loading),
      tap(res => this.error = res.error),
      map(res => res.data)
    );
  }

  rateClicked(id: number): void {
    console.log('rateClicked', id)
  }

  likeClicked(id: number): void {
    console.log('likeClicked', id)
  }

  dislikeClicked(id: number): void {
    console.log('dislikeClicked', id)
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.update();
  }
}
