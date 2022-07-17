import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { InspectorComponent } from "./components/inspector/inspector.component";
import { MatchComponent } from "./components/match/match.component";

const routes: Routes = [
  {path: '', redirectTo: 'inspector', pathMatch: 'full'},
  {path: 'inspector', component: InspectorComponent},
  {path: 'match', component: MatchComponent},
  {path: '**', redirectTo: 'inspector', pathMatch: 'full'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
