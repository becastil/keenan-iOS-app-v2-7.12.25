package com.uber.sydneyhealth.ribs.root

import android.content.Context
import android.util.AttributeSet
import android.widget.FrameLayout

class RootView @JvmOverloads constructor(
    context: Context,
    attrs: AttributeSet? = null,
    defStyle: Int = 0
) : FrameLayout(context, attrs, defStyle), RootInteractor.RootPresenter