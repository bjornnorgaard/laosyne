import { Component, OnInit } from '@angular/core';

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

  constructor() {
  }

  ngOnInit(): void {
  }

}

interface Route {
  route: string;
  icon: string;
}
