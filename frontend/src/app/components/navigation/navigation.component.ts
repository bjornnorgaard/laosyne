import { Component } from '@angular/core';
import { RescanPathsGQL } from "../../../generated/graphql";

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent  {

  public routes: Route[] = [
    {route: 'inspector', icon: 'search'},
    {route: 'match', icon: 'play_arrow'},
  ];

  constructor(private scan: RescanPathsGQL) {
  }

  public reScan(): void {
    this.scan.mutate().subscribe();
  }

}

interface Route {
  route: string;
  icon: string;
}
