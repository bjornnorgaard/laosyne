import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { environment } from "../../../environments/environment";
import {
  DislikePictureGQL,
  LikePictureGQL,
  Picture,
  PictureDetailsGQL,
  RatePictureGQL
} from "../../../generated/graphql";
import { filter, map, Observable, of, tap } from "rxjs";

@Component({
  selector: 'image-details',
  templateUrl: './image-details.component.html',
  styleUrls: ['./image-details.component.scss']
})
export class ImageDetails implements OnChanges {
  @Input() public id: number = 4;
  @Input() public full: boolean = false;
  @Input() public watch: boolean = false;

  public api: string = environment.api;
  public loading: boolean = true;
  public error: any = null;
  public isVideo: boolean = false;
  public picture: Observable<Picture> = of({} as Picture);

  constructor(private query: PictureDetailsGQL,
              private rate: RatePictureGQL,
              private like: LikePictureGQL,
              private dislike: DislikePictureGQL) {
    this.refresh();
  }

  private refresh() {
    let watch = this.query.watch({id: this.id});
    this.picture = watch.valueChanges.pipe(
      tap(res => this.loading = res.loading),
      tap(res => this.error = res.error),
      filter(res => res.data.Picture.__typename != undefined),
      map(res => res.data.Picture),
      tap(res => res.ext === '.webm' ? this.isVideo = true : this.isVideo = false)
    );

    if (this.watch) {
      watch.startPolling(5000);
    }
  }

  rateClicked(id: number): void {
    console.log('rateClicked', id)
    this.rate.mutate({id: id}).subscribe(pic => {
      if (!pic.data) {
        return;
      }
      this.picture = of(pic.data.AddToRating)
    });
  }

  likeClicked(id: number): void {
    console.log('likeClicked', id)
    this.like.mutate({id: id}).subscribe(pic => {
      if (!pic.data) {
        return;
      }
      this.picture = of(pic.data.LikePicture)
    });
  }

  dislikeClicked(id: number): void {
    console.log('dislikeClicked', id)
    this.dislike.mutate({id: id}).subscribe(pic => {
      if (!pic.data) {
        return;
      }
      this.picture = of(pic.data.DislikePicture)
    });
  }

  ngOnChanges(changes: SimpleChanges): void {
    this.refresh();
  }
}
