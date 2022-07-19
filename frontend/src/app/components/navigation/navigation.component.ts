import { Component, OnInit } from '@angular/core';
import { RescanPathsGQL } from "../../../generated/graphql";

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {

  public routes: Route[] = [
    {route: 'inspector', icon: 'search'},
    {route: 'match', icon: 'play_arrow'},
  ];

  constructor(private scan: RescanPathsGQL) {
  }

  ngOnInit(): void {
    this.reScan();
  }

  public reScan(): void {
    this.scan.mutate().subscribe();
  }

}

interface Route {
  route: string;
  icon: string;
}
