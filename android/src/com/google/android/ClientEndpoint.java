package com.google.android;

import android.accounts.AccountManager;
import android.app.Activity;
import android.app.Dialog;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.ListView;
import android.widget.TextView;

/*
import com.google.android.gms.common.GooglePlayServicesUtil;
import com.google.api.client.googleapis.extensions.android.gms.auth.GoogleAccountCredential;
import com.google.api.services.tictactoe.Tictactoe;
import com.google.api.services.tictactoe.model.Board;
import com.google.api.services.tictactoe.model.Score;
import com.google.api.services.tictactoe.model.ScoreCollection;
*/

import java.io.IOException;
import java.util.ArrayList;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * Main activity for the application, it handles the game UI and auth and
 * spawns tasks to Endpoints.
 */
public class ClientEndpoint extends Activity {
  private static final Level LOGGING_LEVEL = Level.ALL;

  private static final String TAG = "TicTacToeSample";

  static final String PREF_ACCOUNT_NAME = "accountName";
  static final int REQUEST_GOOGLE_PLAY_SERVICES = 0;
  static final int REQUEST_ACCOUNT_PICKER = 1;
  static final String PREF_AUTH_TOKEN = "authToken";


  SharedPreferences settings;
  String accountName;

  //GoogleAccountCredential credential;

  boolean signedIn = false;
  boolean waitingForMove = false;

/*
  public void signIn(View v) {
    if (!this.signedIn) {
      chooseAccount();
    } else {
      forgetAccount();
      setSignInEnablement(true);
      setBoardEnablement(false);
      setAccountLabel("(not signed in)");
    }
  }

  private void chooseAccount() {
    startActivityForResult(credential.newChooseAccountIntent(), REQUEST_ACCOUNT_PICKER);
  }

  private void setAccountName(String accountName) {
    SharedPreferences.Editor editor = settings.edit();
    editor.putString(PREF_ACCOUNT_NAME, accountName);
    editor.commit();
    credential.setSelectedAccountName(accountName);
    this.accountName = accountName;
  }

  private void onSignIn() {
    this.signedIn = true;
    this.waitingForMove = true;
    setSignInEnablement(false);
    setBoardEnablement(true);
    setAccountLabel(this.accountName);
    queryScores();
  }

  private void forgetAccount() {
    this.signedIn = false;
    SharedPreferences.Editor editor2 = settings.edit();
    editor2.remove(PREF_AUTH_TOKEN);
    editor2.commit();
  }
*/

  @Override
  public void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.main);

/*
    settings = getSharedPreferences(TAG, 0);
    credential = GoogleAccountCredential.usingAudience(this, ClientCredentials.AUDIENCE);
    setAccountName(settings.getString(PREF_ACCOUNT_NAME, null));

    Tictactoe.Builder builder = new Tictactoe.Builder(
        AndroidHttp.newCompatibleTransport(),
        new GsonFactory(),
        credential);
    service = builder.build();


    if (credential.getSelectedAccountName() != null) {
      onSignIn();
    }
*/
    Logger.getLogger("com.com.google.api.client").setLevel(LOGGING_LEVEL);
  }

}
