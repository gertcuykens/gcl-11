package my.endpoints;
/*
import android.app.Activity;
import android.app.AlertDialog;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.util.Pair;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;
import com.appspot.util.IabHelper;
import com.appspot.util.IabResult;
import com.appspot.util.Inventory;
import com.appspot.util.Purchase;

public class ProductActivity extends Activity implements View.OnClickListener {
    String base64EncodedPublicKey ="";
    IabHelper mHelper;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.products);

        Log.d("IAB", "Creating IAB helper.");
        mHelper = new IabHelper(this, base64EncodedPublicKey);
        mHelper.enableDebugLogging(true);
        Log.d("IAB", "Starting setup.");
        mHelper.startSetup(new IabHelper.OnIabSetupFinishedListener() {
            public void onIabSetupFinished(IabResult result) {
                Log.d("IAB", "Setup finished.");
                mHelper.queryInventoryAsync(mGotInventoryListener);
                if (!result.isSuccess()) { complain("Problem setting up in-app billing: " + result); return;}
                if (mHelper == null) return;
            }
        });

        Button product1Button = (Button) findViewById(R.id.product1Button);
        product1Button.setOnClickListener(this);

        Button product2Button = (Button) findViewById(R.id.product2Button);
        product2Button.setOnClickListener(this);

        Button product3Button = (Button) findViewById(R.id.product3Button);
        product3Button.setOnClickListener(this);
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        Log.d("IAB", "onActivityResult(" + requestCode + "," + resultCode + "," + data);
        if (mHelper == null) return;
        if (!mHelper.handleActivityResult(requestCode, resultCode, data)) {super.onActivityResult(requestCode, resultCode, data);}
        else {Log.d("IAB", "onActivityResult handled by IABUtil.");}
    }

    @Override
    public void onClick(View view) {
        //Context context = view.getContext();
        switch(view.getId()) {
            case R.id.product1Button: mHelper.launchPurchaseFlow(this, "gas", 10001, mPurchaseFinishedListener, "payload"); break;
            case R.id.product2Button: mHelper.launchPurchaseFlow(this, "premium", 10001, mPurchaseFinishedListener, "payload"); break;
            case R.id.product3Button: mHelper.launchPurchaseFlow(this, "infinite_gas", IabHelper.ITEM_TYPE_SUBS, 10001, mPurchaseFinishedListener, "payload"); break;
        }
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        Log.d("IAB", "Destroying helper.");
        if (mHelper != null) {mHelper.dispose(); mHelper = null;}
    }

    IabHelper.QueryInventoryFinishedListener mGotInventoryListener = new IabHelper.QueryInventoryFinishedListener() {
        public void onQueryInventoryFinished(IabResult result, Inventory inventory) {
            Log.d("IAB", "Query inventory finished.");
            if (mHelper == null) return;
            if (result.isFailure()) {
                //complain(c, "Failed to query inventory: " + result);
                return;
            }
            Log.d("IAB", "Query inventory was successful.");
        }
    };

    IabHelper.OnIabPurchaseFinishedListener mPurchaseFinishedListener = new IabHelper.OnIabPurchaseFinishedListener() {
        public void onIabPurchaseFinished(IabResult result, Purchase purchase) {
            Log.d("IAB", "Purchase finished: " + result + ", purchase: " + purchase);
            if (mHelper == null) return;
            if (result.isFailure()) {
                //complain(c, "Error purchasing: " + result); return;
            }
            //if (!verifyDeveloperPayload(purchase)) { MainActivity.complain("Error purchasing. Authenticity verification failed."); return; }
            Log.d("IAB", "Purchase successful.");
            //purchase.getSku().equals(SKU_GAS)
        }
    };

    IabHelper.OnConsumeFinishedListener mConsumeFinishedListener = new IabHelper.OnConsumeFinishedListener() {
        public void onConsumeFinished(Purchase purchase, IabResult result) {
            Log.d("IAB", "Consumption finished. Purchase: " + purchase + ", result: " + result);
            if (mHelper == null) return;
            if (result.isSuccess()) { Log.d("IAB", "Consumption successful. Provisioning.");}
            else {
                //complain(c, "Error while consuming: " + result);
            }
            Log.d("IAB", "End consumption flow.");
        }
    };

    private void complain(String message) {
        Log.e("IAB", "**** Error: " + message);
        AlertDialog.Builder bld = new AlertDialog.Builder(this);
        bld.setMessage(message);
        bld.setNeutralButton("OK", null);
        Log.d("IAB", "Showing alert dialog: " + message);
        bld.create().show();
    }
}
*/

/*
Purchase premiumPurchase = inventory.getPurchase(SKU_PREMIUM);
mIsPremium = (premiumPurchase != null && verifyDeveloperPayload(premiumPurchase));
Log.d("IAB", "User is " + (mIsPremium ? "PREMIUM" : "NOT PREMIUM"));

Purchase infiniteGasPurchase = inventory.getPurchase(SKU_INFINITE_GAS);
mSubscribedToInfiniteGas = (infiniteGasPurchase != null && verifyDeveloperPayload(infiniteGasPurchase));
Log.d("IAB", "User " + (mSubscribedToInfiniteGas ? "HAS" : "DOES NOT HAVE") + " infinite gas subscription.");
if (mSubscribedToInfiniteGas) mTank = TANK_MAX;

// Check for gas delivery -- if we own gas, we should fill up the tank immediately
Purchase gasPurchase = inventory.getPurchase(SKU_GAS);
if (gasPurchase != null && verifyDeveloperPayload(gasPurchase)) {
    Log.d("IAB", "We have gas. Consuming it.");
    mHelper.consumeAsync(inventory.getPurchase(SKU_GAS), mConsumeFinishedListener);
    return;
}
Log.d("IAB", "Initial inventory query finished; enabling main UI.");
*/

/*
    void saveData() {
        SharedPreferences.Editor spe = getPreferences(MODE_PRIVATE).edit();
        spe.putInt("tank", mTank);
        spe.commit();
        Log.d("IAB", "Saved data: tank = " + String.valueOf(mTank));
    }

    void loadData() {
        SharedPreferences sp = getPreferences(MODE_PRIVATE);
        mTank = sp.getInt("tank", 2);
        Log.d("IAB", "Loaded data: tank = " + String.valueOf(mTank));
    }
*/