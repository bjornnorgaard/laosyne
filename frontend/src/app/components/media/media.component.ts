import { Component, Input } from '@angular/core';
import { environment } from "../../../environments/environment";

@Component({
  selector: 'app-media',
  templateUrl: './media.component.html',
  styleUrls: ['./media.component.scss']
})
export class MediaComponent {
  @Input() public id: number = 6;
  @Input() public full: boolean = false;

  public api: string = environment.api;

  rateClicked(id: number): void {

  }

  likeClicked(id: number): void {

  }

  dislikeClicked(id: number): void {

  }
}
