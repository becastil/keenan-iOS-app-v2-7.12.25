package com.uber.sydneyhealth

import android.os.Bundle
import android.view.ViewGroup
import com.uber.rib.core.RibActivity
import com.uber.rib.core.ViewRouter
import com.uber.sydneyhealth.ribs.root.RootBuilder
import com.uber.sydneyhealth.ribs.root.RootRouter

class RootActivity : RibActivity() {
    
    override fun createRouter(parentViewGroup: ViewGroup): ViewRouter<*, *> {
        val rootBuilder = RootBuilder(object : RootBuilder.ParentComponent {})
        return rootBuilder.build(parentViewGroup)
    }
    
    override fun onCreate(savedInstanceState: Bundle?) {
        setTheme(R.style.Theme_SydneyHealth)
        super.onCreate(savedInstanceState)
    }
    
    override fun onBackPressed() {
        val rootRouter = router as? RootRouter
        if (rootRouter?.handleBackPress() != true) {
            super.onBackPressed()
        }
    }
}