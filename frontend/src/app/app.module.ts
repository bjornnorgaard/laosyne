import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GraphQLModule } from './graphql.module';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from "./material/material.module";
import { ImageDetails } from './components/image-details/image-details.component';
import { NavigationComponent } from './components/navigation/navigation.component';
import { MatchComponent } from './components/match/match.component';
import { InspectorComponent } from './components/inspector/inspector.component';

@NgModule({
  declarations: [
    AppComponent,
    ImageDetails,
    NavigationComponent,
    MatchComponent,
    InspectorComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    GraphQLModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MaterialModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
